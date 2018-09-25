<template>
    <div>
        <el-table :data="data" style="width: 100%" stripe height="600"
                  element-loading-text="拼命加载中"
                  element-loading-spinner="el-icon-loading"
                  element-loading-background="rgba(0, 0, 0, 0.8)"
                  v-loading="tableLoading"
        >
            <el-table-column prop="id" label="标志"></el-table-column>
            <el-table-column prop="user" label="账号" width="100"></el-table-column>
            <el-table-column prop="password" label="密码" width="100"></el-table-column>
            <el-table-column prop="dbname" label="实例名" width="199"></el-table-column>
            <el-table-column prop="host" label="主机" width="400"></el-table-column>
            <el-table-column prop="port" label="端口" width="400"></el-table-column>
        </el-table>
        <div class="block">
            <el-pagination
                    layout="prev, pager, next"
                    @current-change="pg_change"
                    :page-count="totalPg">
            </el-pagination>
        </div>

    </div>
</template>

<script>
    import Api from "../api"
    export default {
        name:"DbInfoList",
        data() {
            return {
                totalPg: 0,
                currentPg: 1,
                tableLoading:false,
                data: []
            }
        },methods:{
            pg_change: function (pg) {
                this.tableLoading = true;
                this.currentPg = pg;
                var that = this;
                var param = {
                    size: 15,
                    pg: that.currentPg,
                };
                Api.list(param, function (data) {
                    that.tableLoading = false;
                    that.totalPg =  data.totalPg;
                    that.data = data.data;
                })
            }
        },mounted: function () {
            var that=this;
            this.$nextTick(function () {
                that.pg_change(1);
            })
        },
        watch: {
            data (val) {
                this.data = val;
            }
        }
    }
</script>