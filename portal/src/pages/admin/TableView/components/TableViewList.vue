<template>
        <el-table
                :data="columnData"
                :height="500"
                :span-method="objectSpanMethod"
                border
                style="width: 100%;">
            <el-table-column
                    prop="relname"
                    label="表信息"
                    width="180">
                <template slot-scope="scope">
                    <el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <span>{{ scope.row.relname}}</span>
                        </div>
                        <div class="text item">
                            <el-button style="float: right; padding: 3px 0" @click="commentTable(scope.row)" type="mini"
                                       icon="el-icon-edit"></el-button>
                            {{ scope.row.tableinfo.description}}
                        </div>
                    </el-card>
                </template>
            </el-table-column>
            <el-table-column
                    prop="attname"
                    width="150px"
                    label="列名">
            </el-table-column>
            <el-table-column
                    prop="description"
                    label="说明">
                <template slot-scope="scope">
                    <div class="text item">
                        {{ scope.row.description }}
                        <el-button style="float: right; padding: 3px 0" @click="commentColumn(scope.row)" type="mini"
                                   icon="el-icon-edit"></el-button>
                    </div>
                </template>
            </el-table-column>

            <el-table-column
                    prop="typname"
                    width="80px"
                    label="数据类型">
            </el-table-column>
            <el-table-column
                    prop="attlen"
                    width="80px"
                    label="数据长度">
                <template slot-scope="scope">
                    <el-tag type="info" v-if="scope.row.attlen ==-1 ">
                        不限
                    </el-tag>
                    <el-tag type="success" v-else>
                        {{ scope.row.attlen }}
                    </el-tag>
                </template>
            </el-table-column>
            <el-table-column
                    prop="attnotnull"
                    width="80px"
                    label="必填">
                <template slot-scope="scope">
                        <el-tag type="success" v-if="scope.row.attnotnull ">
                            是
                        </el-tag>
                        <el-tag type="info" v-else>
                            否
                        </el-tag>
                </template>
            </el-table-column>
            <el-table-column
                    prop="adsrc"
                    width="300px"
                    label="默认值">
            </el-table-column>
            <el-table-column
                    width="80px"
                    prop="attisdropped"
                    label="已删">
                <template slot-scope="scope">
                        <el-tag type="info"  v-if="scope.row.attisdropped ">
                            是
                        </el-tag>
                        <el-tag type="success" v-else>
                            否
                        </el-tag>
                </template>
            </el-table-column>
        </el-table>
</template>

<script>
    import Api from "./api"
    import {mapMutations} from "vuex"
    export default {
        name: "TableViewList",
        data() {
            return {
                columnData: [],
            }
        },
        mounted: function () {
            var that = this;
            this.$nextTick(function () {
                that.loadColumns();
            })
        },
        methods: {
            commentTable: function (row) {
                this.SET_MDF_RESOURCE(row);
                this.SET_MDF_TABLE_VISIBLE(true);
            }, commentColumn: function (row) {
                this.SET_MDF_RESOURCE(row);
                this.SET_MDF_COLUMN_VISIBLE(true);
            },
            refresh: function (id) {
                this.loadColumns(id);
            },
            loadColumns: function (id) {
                var that = this;
                Api.columns({id: id}, function (res) {
                    console.log(id);
                    that.columnData = res;
                });
            },
            objectSpanMethod({row, column, rowIndex, columnIndex}) {
                var that = this;
                if (columnIndex === 0) {
                    var diffFlag = false;
                    let nowVal = that.columnData[rowIndex].relname;
                    if (rowIndex == 0) {
                        diffFlag = true;
                    } else {
                        let preVal = that.columnData[rowIndex - 1].relname;
                        diffFlag = (nowVal != preVal);
                    }
                    if (!diffFlag) {
                        return {
                            rowspan: 0,
                            colspan: 1
                        };
                    }
                    var hasNext = (that.columnData.length > rowIndex);
                    if (!hasNext) {
                        return {
                            rowspan: 1,
                            colspan: 1
                        };
                    }
                    var rowCount = 1;
                    for (var i = rowIndex + 1; i < that.columnData.length; i++) {
                        let val = that.columnData[i].relname;
                        if (val == nowVal) {
                            rowCount++;
                        } else {
                            return {
                                rowspan: rowCount,
                                colspan: 1
                            };
                        }
                    }
                    return {
                        rowspan: rowCount,
                        colspan: 1
                    };
                }

                return {
                    rowspan: 1,
                    colspan: 1
                };
            },
            ...mapMutations('admin/tableview', ['SET_MDF_TABLE_VISIBLE', 'SET_MDF_COLUMN_VISIBLE', "SET_MDF_RESOURCE"])
        },
        watch: {
            data (val) {
                this.data = val;
            },
            columns(val){
                this.columns = val;
            }
        }
    };
</script>
