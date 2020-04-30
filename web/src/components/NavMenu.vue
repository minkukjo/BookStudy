<template>
  <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect" background-color="#34495e" text-color="#fff" active-text-color="#ffd04b" style="border-bottom-color: #2c3e50">
    <el-menu-item index="1" class="el-icon-s-home"></el-menu-item>
    <el-submenu index="2">
      <template slot="title">Board</template>
      <el-menu-item index="2-1">자유 게시판</el-menu-item>
      <el-menu-item index="2-2">스터디 모집</el-menu-item>
      <el-menu-item index="2-3">정보 게시판</el-menu-item>
    </el-submenu>
    <el-submenu index="3" style="float: right">
      <template slot="title"><i class="el-icon-user"></i></template>
      <el-menu-item index="3-1">My Account</el-menu-item>
      <el-menu-item index="3-2">Logout</el-menu-item>
    </el-submenu>
  </el-menu>
</template>

<script>
export default {
  name: 'NavMenu',
  data: function () {
    return {
      info: null,
      activeIndex: '1'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      this.activeIndex = key
      this.$emit('prop', this.activeIndex)
      console.log(key)
    }
  }
}
</script>

<style>

</style>
