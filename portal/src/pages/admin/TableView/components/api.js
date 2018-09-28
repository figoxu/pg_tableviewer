import axios from "axios";

const Api = {
    columns:function (data,callback) {
        var url = "/api/pg_db_info/table_info/"+data.id+"/columns";
        axios({
            method: "GET",
            url: url,
        }).then(function (res) {
            console.log("xxxx")
            console.log(res)
            if (callback) callback(res);
        }).catch(function (error) {
        });
    }
};

export default Api;
