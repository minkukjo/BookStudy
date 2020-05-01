<template>
  <div class="qna">
    <NavMenu :prop-index="activeIndex"></NavMenu>
    <Title :prop-title="title"></Title>
    <Board></Board>
  </div>
</template>

<script>
export default {
  name: 'Qna',
  data: function () {
    return {
      activeIndex: '2-3',
      title: '질문 게시판'
    }
  }
}
</script>

<style scoped>

</style>
