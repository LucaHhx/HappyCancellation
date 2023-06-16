<template>
    <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>积分：{{integration}}</span>
        <el-button type="primary" size="default" @click="refresh">刷新游戏</el-button>
        
      </div>
    </template>
   <div id="div1">
    <el-row v-for="(row, i) in table" :key="i">
      <el-col class="col" span="1" v-for="(col, ii) in row" :key="i + '|' + ii">
        <el-button
          :class="'but iconfont  icon-'+col.Name"
          :type="types[col.Status]"
          size="mini"
          :key="i + '-' + ii"
          @click="click(col)"
          :id="col.Uid"
          :style="col.Style"
        >
        </el-button>
      </el-col>
    </el-row>
  </div></el-card>
</template>
<script setup>
import { ElMessage } from "element-plus";
import { ref } from "vue";
var types = ["primary", "success", "warning", "danger", "info"];
var table = ref([]);
var integration = ref(0)
const onNewGame = (data) => {
  table.value = data.Table;
  integration.value = data.Integration
  console.log(data);
};
const onError = (data) => {
  console.log(data);
  ElMessage({
    message: data,
    type: "error",
    duration: 3 * 1000,
  });
};
const refresh = () => {
  window.starx.request("Serve.NewGame", { name: "huang", id: 1,refresh:true }, onNewGame);
};
const gateHost = "47.99.106.90";
const gatePort = 4321;
window.starx.init(
  { host: gateHost, port: gatePort, path: "/nano" },
  function () {
    window.starx.on("onNewGame", onNewGame);
    window.starx.on("onError", onError);
    window.starx.request("Serve.NewGame", { name: "huang", id: 1 }, onNewGame);
  }
);
var tagOne = "";
const click = (col) => {
  if (tagOne == "") {
    tagOne = col;
    col.Status = 0
  } else {
    window.starx.notify("Serve.PositionExchange", {
      user: { name: "huang", id: 1 },
      tag1: tagOne,
      tag2: col,
    });
    tagOne = "";
  }
};
</script>
<style>
.el-button.but {
  width: 50px;
  height: 50px;
  margin: 0px;
  padding: 0px;
  border: 1px solid #ccc;
  position: relative;
  float: left;
  /* transition: 1s all ease 0s; */
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
