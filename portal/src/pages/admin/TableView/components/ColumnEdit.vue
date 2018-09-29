<template>
    <div>
        <el-form :model="form" :label-position="labelPosition" :inline="inline" >
            <el-form-item label="表名">
                <el-input v-model="form.relname" :readonly="true"></el-input>
            </el-form-item>
            <el-form-item label="表描述">
                <el-input v-model="form.tableinfo.description" ></el-input>
            </el-form-item>
            <el-form-item label="列名">
                <el-input v-model="form.attname" readonly="true"></el-input>
            </el-form-item>
            <el-form-item label="列描述">
                <el-input v-model="form.description" ></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="SET_MDF_COLUMN_VISIBLE(false)">取 消</el-button>
            <el-button type="primary" @click="modify()">确 定</el-button>
        </div>
    </div>
</template>
<script>
    import Api from "./api"
    import {mapMutations} from "vuex"
    export default {
        name: "ColumnEdit",
        data() {
            return {
                inline:true,
                labelPosition:"left",
            }
        },
        computed: {
            form:{
                get: function () {
                    return this.$store.state.admin.tableview.mdfResource;
                },
                set: function (newValue) {
                    this.SET_MDF_RESOURCE( newValue );
                }
            }
        },methods:{
            modify:function(){
                var form = this.form;
                Api.commentColumn({
                    id: form.dbid,
                    comment: form.description,
                    attname: form.attname,
                    relname: form.relname,
                    nspname: form.nspname,
                });
                this.SET_MDF_COLUMN_VISIBLE(false)
            },
            ...mapMutations('admin/tableview', ['SET_MDF_TABLE_VISIBLE', 'SET_MDF_COLUMN_VISIBLE',"SET_MDF_RESOURCE"])
        }
    }
</script>