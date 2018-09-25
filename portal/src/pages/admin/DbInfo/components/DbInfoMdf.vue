<template lang="html">
    <div>
        <el-form :model="form" :label-position="labelPosition" :inline="inline" >
            <el-form-item label="账号">
              <el-input v-model="form.user" ></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="form.password" ></el-input>
            </el-form-item>
            <el-form-item label="实例名">
              <el-input v-model="form.dbname" ></el-input>
            </el-form-item>
            <el-form-item label="主机">
              <el-input v-model="form.host" ></el-input>
            </el-form-item>
            <el-form-item label="端口">
              <el-input v-model="form.port" ></el-input>
            </el-form-item>
         </el-form>
        <div slot="footer" class="dialog-footer">
           <el-button @click="SET_MDF_VISIBLE(false)">取 消</el-button>
           <el-button type="primary" @click="modify()">确 定</el-button>
        </div>
    </div>
</template>

<script>

    import Api from "./api"
    import { mapMutations } from "vuex"

    export default {
        name: 'DbInfoMdf',
        data() {
            return {
                inline:true,
                labelPosition:"left",
            }
        },
        computed: {
            form:{
                get: function () {
                    return this.$store.state.admin.dbinfo.mdfResource;
                },
                set: function (newValue) {
                    this.SET_MDF_RESOURCE( newValue );
                }
            }
        },methods:{
            modify:function(){
                var that = this;
                Api.mdf(this.form,function () {
                    that.SET_MDF_VISIBLE(false)
                    that.$emit("done",true);
                })
            },
            ...mapMutations('admin/dbinfo', ['SET_MDF_VISIBLE',"SET_MDF_RESOURCE"])
        }
    }
</script>