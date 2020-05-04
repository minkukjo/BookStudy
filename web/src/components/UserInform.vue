<template>
  <el-card shadow="always" class="box-card" style="margin-top: 30px;">
    <div v-for="o in this.userData" :key="o" class="text item">
      {{ o }}
    </div>
  </el-card>
</template>

<script>
export default {
  name: 'UserInform',
  props: {
    propUserData: Object
  },
  data: function () {
    return {
      userData: this.propUserData
    }
  }
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
