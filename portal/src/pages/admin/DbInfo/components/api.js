import axios from "axios";

const Api = {
    add(data,callback){
        axios({
            method: "POST",
            url: "/api/pg_db_info/add",
            data: data,
        }).then(function (res) {
            if (callback) callback(res.data);
        }).catch(function (error) {
        });
    },
    list(data, callback){
        var url = "/api/pg_db_info/list/" + data.size + "/" + data.pg;
        axios({
            method: "GET",
            url: url,
        }).then(function (res) {
            if (callback) callback(res.data);
        }).catch(function (error) {
        });
    },
    mdf(data,callback){
        var url = "/api/pg_db_info/update";
        axios({
            method: "POST",
            url: url,
        }).then(function (res) {
            if (callback) callback(res.data);
        }).catch(function (error) {
        });

    },
    del(data,callback){
        var url = "/api/pg_db_info/del/"+data.id;
        axios({
            method: "POST",
            url: url,
        }).then(function (res) {
            if (callback) callback(res.data);
        }).catch(function (error) {
        });

    },
};

export default Api;
