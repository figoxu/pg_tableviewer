<template>
    <div>
        <el-row>
            <el-button type="primary" @click="SET_ADD_VISIBLE(true)">添加数据库</el-button>
        </el-row>
        <DbInfoList></DbInfoList>
        <el-dialog title="数据库-增加" :visible.sync="addVisible">
            <DbInfoAdd></DbInfoAdd>
        </el-dialog>

        <el-dialog title="数据库-修改" :visible.sync="mdfVisible">
            <DbInfoMdf></DbInfoMdf>
        </el-dialog>
    </div>
</template>

<script>
    import DbInfoList from './components/DbInfoList'
    import DbInfoAdd from "./components/DbInfoAdd"
    import DbInfoMdf from './components/DbInfoMdf'
    import {mapMutations} from "vuex";

    export default {
        name: 'DbInfo',
        components: {DbInfoList,DbInfoAdd,DbInfoMdf},
        computed: {
            mdfVisible:{
                get: function () {
                    return this.$store.state.admin.dbinfo.mdfVisible;
                },
                set: function (newValue) {
                    this.SET_MDF_VISIBLE( newValue );
                }
            },
            addVisible:{
                get: function () {
                    return this.$store.state.admin.dbinfo.addVisible;
                },
                set: function (newValue) {
                    this.SET_ADD_VISIBLE( newValue );
                }
            }
        },
        methods: {
            ...mapMutations('admin/dbinfo', ['SET_MDF_VISIBLE', 'SET_ADD_VISIBLE'])
        }
    };
</script>
