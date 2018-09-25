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
    import Treeselect from '@riophae/vue-treeselect'
    import '@riophae/vue-treeselect/dist/vue-treeselect.css'
    import {mapMutations, mapState} from 'vuex'

    export default {
        name: 'ResourceAdd',
        components: { Treeselect },
        data() {
            return {
                form: {
                    pid : 0,
                    name: '',
                    sysName : '',
                    priority : 0,
                    path : '',
                    type : '',
                    permission : '',
                    available : true
                },
                options: [],
                inline:true,
                labelPosition:"left",
            }
        },
        computed: {
        },
        methods: {
            commitAdd:function () {
                var data = this.form;
                Api.add(data);
               this.SET_ADD_VISIBLE(false);
            },
            ...mapMutations(['SET_ADD_VISIBLE'])
        },mounted: function () {
            var that=this;
            this.$nextTick(function () {
                that.loadOptions();
            })
        },
        watch: {
            data (val) {
                this.data = val;
            }
        }
    };
</script>