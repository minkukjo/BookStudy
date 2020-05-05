<template>
  <div style="text-align: left; margin-top: 20px; font-weight: bold">
  {{this.title}}
    <div style="float: right">
      <el-button type="info" @click="handlerClick" icon="el-icon-edit"></el-button>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Title',
  props: {
    propTitle: String
  },
  data: function () {
    return {
      title: this.propTitle
    }
  },
  methods: {
    handlerClick: function () {
      this.$router.push('/main/write')
    }
  }
}
</script>

<style scoped>
  .el-button--info {
    background: #34495e;
  }
</style>
