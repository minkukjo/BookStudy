<template>
  <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect" background-color="#34495e" text-color="#fff" active-text-color="#ffd04b" style="border-bottom-color: #2c3e50">
    <el-menu-item index="1" class="el-icon-s-home"></el-menu-item>
    <el-submenu index="2">
      <template slot="title">Board</template>
      <el-menu-item index="2-1">자유 게시판</el-menu-item>
      <el-menu-item index="2-2">스터디 모집</el-menu-item>
      <el-menu-item index="2-3">질문 게시판</el-menu-item>
    </el-submenu>
    <el-submenu index="3" style="float: right">
      <template slot="title"><i class="el-icon-user"></i></template>
      <el-menu-item index="3-1">My Account</el-menu-item>
      <el-menu-item index="3-2"><a href="/logout">Logout</a></el-menu-item>
    </el-submenu>
  </el-menu>
</template>

<script>
export default {
  name: 'NavMenu',
  props: {
    propIndex: String
  },
  data: function () {
    return {
      info: null,
      activeIndex: this.propIndex
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      const oldIndex = this.activeIndex
      this.activeIndex = key

      if (this.activeIndex === oldIndex && this.$route.params.id === undefined) {
        this.$forceUpdate()
      } else if (this.activeIndex === '1') {
        this.$router.push('/main')
      } else if (this.activeIndex === '2-1') {
        this.$router.push('/main/general')
      } else if (this.activeIndex === '2-2') {
        this.$router.push('/main/study')
      } else if (this.activeIndex === '2-3') {
        this.$router.push('/main/qna')
      } else if (this.activeIndex === '3-1') {
        this.$router.push('/main/user')
      }
    }
  }
}
</script>

<style>
  .el-menu-item a {
    display: block;
    text-decoration: none;
    color: white;
  }

  .el-submenu__title i {
    color: white;
  }
</style>
