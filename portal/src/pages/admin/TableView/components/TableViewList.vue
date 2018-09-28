<template>
    <div>
        <el-table
                :data="columnData"
                :height="500"
                border
                style="width: 100%;">
            <el-table-column
                    prop="relname"
                    label="表名"
                    width="180">
            </el-table-column>
            <el-table-column
                    prop="nspname"
                    label="表空间">
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
//                if (columnIndex === 0) {
//                    if (rowIndex % 2 === 0) {
//                        return {
//                            rowspan: 2,
//                            colspan: 1
//                        };
//                    } else {
//                        return {
//                            rowspan: 0,
//                            colspan: 0
//                        };
//                    }
//                }

                return {
                    rowspan: 0,
                    colspan: 0
                };
            }
        },mounted: function () {
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
