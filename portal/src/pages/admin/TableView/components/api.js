import axios from "axios";
import http from "../../../../common/http"

const Api = {
    dbs: function (data, callback) {
        axios({
            method: "GET",
            url: "/api/pg_db_info/all",
        }).then(function (res) {
            console.log("xxxx");
            console.log(res);
            if (callback) callback(res);
        }).catch(function (error) {
        });
    },
    columns: function (data, callback) {
        var url = "/api/pg_db_info/table_info/" + data.id + "/columns";
        axios({
            method: "GET",
            url: url,
        }).then(function (res) {
            console.log("xxxx")
            console.log(res)
            if (callback) callback(res);
        }).catch(function (error) {
        });
    },
    commentTable: function (data, callback) {
        var url = "/api/pg_db_info/table_info/" + data.id + "/comment/table";
        axios({
            method: "PUT",
            url: url,
            transformRequest: [http.transJson2From],
            data: data,
        }).then(function (res) {
            if (callback) callback(res);
        }).catch(function (error) {
        });
    },
    commentColumn: function (data, callback) {
        var url = "/api/pg_db_info/table_info/" + data.id + "/comment/column";
        axios({
            method: "PUT",
            url: url,
            transformRequest: [http.transJson2From],
            data: data,
        }).then(function (res) {
            if (callback) callback(res);
        }).catch(function (error) {
        });
    }
};

export default Api;
