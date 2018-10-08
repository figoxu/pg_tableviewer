<template>
    <div>
        <el-select v-model="dbid" @change="loadList" placeholder="请选择">
            <el-option
                    v-for="item in options"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
            </el-option>
        </el-select>
        <TableViewList  ref="pageList"></TableViewList>
        <el-dialog title="编辑表描述" :visible.sync="mdf_table_visible">
            <TableEdit v-on:done="loadList"></TableEdit>
        </el-dialog>

        <el-dialog title="编辑列描述" :visible.sync="mdf_column_visible">
            <ColumnEdit v-on:done="loadList"></ColumnEdit>
        </el-dialog>
    </div>
</template>


<script>
    import Api from './components/api'
    import TableViewList from './components/TableViewList'
    import TableEdit from './components/TableEdit'
    import ColumnEdit from './components/ColumnEdit'
    import {mapMutations} from "vuex";

    export default {
        name: 'DbInfo',
        components: {TableViewList,TableEdit,ColumnEdit},
        data() {
            return {
                options: [],
                dbid: ''
            }
        },
        computed: {
            mdf_table_visible: {
                get: function () {
                    return this.$store.state.admin.tableview.mdfTableVisible;
                },
                set: function (newValue) {
                    this.SET_MDF_TABLE_VISIBLE(newValue);
                }
            },
            mdf_column_visible: {
                get: function () {
                    return this.$store.state.admin.tableview.mdfColumnVisible;
                },
                set: function (newValue) {
                    this.SET_MDF_COLUMN_VISIBLE(newValue);
                }
            }
        },mounted: function () {
            var that = this;
            this.$nextTick(function () {
                that.loadOptions();
            })
        },
        methods: {
            chgOption:function () {

            },
            loadOptions:function(){
                var that = this;
                Api.dbs({},function (data) {
                    that.options = data;
                })
            },
            loadList: function () {
                var dbid = this.dbid;
                this.$refs.pageList.refresh(dbid);
            },
            ...mapMutations('admin/tableview', ['SET_MDF_TABLE_VISIBLE', 'SET_MDF_COLUMN_VISIBLE'])
        },
        watch: {
            options (val) {
                this.options = val;
            }
        }
    };
</script>
