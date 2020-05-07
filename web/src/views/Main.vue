<template>
  <div class="main">
    <el-collapse-transition>
      <Carousel></Carousel>
    </el-collapse-transition>
  </div>
</template>

<script>

export default {
  data: function () {
    return {
      activeIndex: '1'
    }
  }
}
</script>
