<template>
    <div>
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
    import TableViewList from './components/TableViewList'
    import TableEdit from './components/TableEdit'
    import ColumnEdit from './components/ColumnEdit'
    import {mapMutations} from "vuex";

    export default {
        name: 'DbInfo',
        components: {TableViewList,TableEdit,ColumnEdit},computed: {
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
        },
        methods: {
            loadList: function (v) {
                this.$refs.pageList.refresh();
            },
            ...mapMutations('admin/tableview', ['SET_MDF_TABLE_VISIBLE', 'SET_MDF_COLUMN_VISIBLE'])
        }
    };
</script>
