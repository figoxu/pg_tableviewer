<template>
    <div>
        <el-table
                :data="columnData"
                :height="500"
                :span-method="objectSpanMethod"
                border
                style="width: 100%;">
            <el-table-column
                    prop="relname"
                    label="表名"
                    width="180">
            </el-table-column>
            <el-table-column
                    prop="tableinfo.description"
                    label="表说明">
            </el-table-column>
            <el-table-column
                    prop="attname"
                    label="列名">
            </el-table-column>
            <el-table-column
                    prop="description"
                    label="说明">
            </el-table-column>
        </el-table>
    </div>
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
            loadColumns: function () {
                var that = this;
                Api.columns({id: 11}, function (res) {
                    console.log(res);
                    that.columnData = res;
                });
            },
            objectSpanMethod({row, column, rowIndex, columnIndex}) {
                var that = this;
                if (columnIndex === 0 || columnIndex===1) {
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
            }
        }, mounted: function () {
            var that = this;
            this.$nextTick(function () {
                that.loadColumns();
            })
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
