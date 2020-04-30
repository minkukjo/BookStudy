<template>
  <div class="main">
    <NavMenu @prop="indexHandler"></NavMenu>
    <el-collapse-transition>
      <Carousel v-if="activeIndex === '1'"></Carousel>
      <Board v-if="activeIndex === '2-1'" ></Board>
      <MyAccount v-if="activeIndex === '3-1'"></MyAccount>
    </el-collapse-transition>
  </div>
</template>

<script>

export default {
  el: '#main',
  data: function () {
    return {
      activeIndex: '1'
    }
  },
  methods: {
    indexHandler: function (that) {
      this.activeIndex = that
    }
  }
}
</script>
