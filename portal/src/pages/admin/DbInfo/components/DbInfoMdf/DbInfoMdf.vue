<template lang="html">
    <div>
        <el-form :model="form" :label-position="labelPosition" :inline="inline" >
            <el-form-item label="用户名">
              <el-input v-model="form.name" readonly></el-input>
            </el-form-item>
            <el-form-item label="用户分组">
                <el-select v-model="form.gids" multiple placeholder="请选择" style="width: 100%;">
                <el-option
                  v-for="item in options"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id">
                </el-option>
              </el-select>
            </el-form-item>
         </el-form>
        <div slot="footer" class="dialog-footer">
           <el-button @click="SET_MDF_VISIBLE(false)">取 消</el-button>
           <el-button type="primary" @click="modify()">确 定</el-button>
        </div>
    </div>
</template>

<script>

    import Api from "../api"
    import { mapMutations } from "vuex"

    export default {
        name: 'DbInfoMdf',
        data() {
            return {
                options: [],
                inline:true,
                labelPosition:"left",
            }
        },
        computed: {
            form:{
                get: function () {
                    return this.$store.state.user.mdfResource;
                },
                set: function (newValue) {
                    this.SET_MDF_RESOURCE( newValue );
                }
            }
        },methods:{
            loadOptions:function(){
                var that = this;
                Api.all({},function(data){
                    that.options = data;
                })
            },modify:function(){
                var that = this;
                Api.mdf(this.form,function () {
                    that.SET_MDF_VISIBLE(false)
                })
            },
            ...mapMutations(['SET_MDF_VISIBLE',"SET_MDF_RESOURCE"])
        },mounted: function () {
            var that=this;
            this.$nextTick(function () {
                that.loadOptions();
            })
        }

    }
</script>