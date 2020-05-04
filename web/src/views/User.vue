<template>
  <div class="userInform">
    <NavMenu :prop-index="activeIndex"></NavMenu>
    <el-collapse-transition v-if="userData !== null">
      <UserInform :prop-user-data="userData" ></UserInform>
    </el-collapse-transition>
  </div>
</template>

<script>

export default {
  name: 'User',
  created () {
    this.getUserInform()
  },
  data: function () {
    return {
      activeIndex: '3-1',
      userData: null
    }
  },
  methods: {
    getUserInform: function () {
      this.$http.get(this.$http_url + '/api/user')
        .then((response) => {
          console.log(response.data)
          console.log(response.data.id)
          console.log(response.data.connected_at)
          console.log(response.data.nickname)
          console.log(response.data.token)
          this.userData = response.data
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>

<style scoped>

</style>
