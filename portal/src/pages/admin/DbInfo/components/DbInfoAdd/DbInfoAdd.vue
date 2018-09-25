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
           <el-button @click="SET_ADD_VISIBLE(false)">取 消</el-button>
           <el-button type="primary" @click="commitAdd()">确 定</el-button>
        </div>
    </div>
</template>

<script>
    import Api from "../api"
    import {mapMutations, mapState} from 'vuex'

    export default {
        name: 'DbInfoAdd',
        data() {
            return {
                form: {
                    user : '',
                    password: '',
                    dbname : '',
                    host : '',
                    port : ''
                },
                inline:true,
                labelPosition:"left",
            }
        },
        computed: {
        },
        methods: {
            commitAdd:function () {
                var that = this;
                Api.add(this.form,function () {
                    that.SET_ADD_VISIBLE(false);
                    that.$emit("done",true);
                });
            },
            ...mapMutations('admin/dbinfo', [
                'SET_ADD_VISIBLE'
            ])
        },
        watch: {
            data (val) {
                this.data = val;
            }
        }
    };
</script>