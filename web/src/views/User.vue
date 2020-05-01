<template>
  <div class="userInform">
    <NavMenu :prop-index="activeIndex"></NavMenu>
    <el-collapse-transition>
    <UserInform></UserInform>
    </el-collapse-transition>
  </div>
</template>

<script>
export default {
  name: 'User',
  data: function () {
    return {
      activeIndex: '3-1'
    }
  }
}
</script>

<style scoped>

</style>
