<template>
  <el-card shadow="always" class="box-card" style="margin-top: 30px;">
    <div v-for="o in 4" :key="o" class="text item">
      {{'List item ' + o }}
    </div>
  </el-card>
</template>

<script>
export default {
  name: 'MyAccount'
}
</script>

<style>
  .text {
    font-size: 14px;
    color: white;
  }

  .item {
    padding: 18px 0;
  }

  .box-card {
    width: 480px;
    margin: 0 auto;
    background-color: #34495e;
  }
</style>
