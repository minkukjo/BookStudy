<template>
  <div>
  <el-menu class="el-menu-demo" mode="horizontal" @select="handleSelect" background-color="#34495e" text-color="#fff" active-text-color="#ffd04b" style="border-bottom-color: #2c3e50">
    <el-menu-item index="/" class="el-icon-s-home">
    </el-menu-item>
    <el-submenu index="2">
      <template slot="title">Board</template>
      <el-menu-item index="/general">자유 게시판</el-menu-item>
      <el-menu-item index="/study">스터디 모집</el-menu-item>
      <el-menu-item index="/qna">질문 게시판</el-menu-item>
    </el-submenu>
    <el-submenu index="3" style="float: right">
      <template slot="title"><i class="el-icon-user"></i></template>
      <el-menu-item index="/user">My Account</el-menu-item>
      <el-menu-item index="logout"><a href="/logout">Logout</a></el-menu-item>
    </el-submenu>
  </el-menu>
  <div id="app">
    <router-view/>
  </div>
  </div>
</template>

<script>
export default {
  methods: {
    handleSelect (index) {
      if (index !== 'logout') {
        this.$router.push(index).catch(err => {
          if (err.name !== 'NavigationDuplicated') {
            throw err
          }
        })
      }
    }
  }
}
</script>

<style>
body {
  background-color:#2c3e50;
}
a {
  text-decoration: none;
}
.el-menu-item a {
    display: block;
    text-decoration: none;
    color: white;
}

.el-submenu__title i {
    color: white !important;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #ecf0f1;
}
</style>
