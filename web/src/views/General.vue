<template>
  <div class="general">
    <NavMenu :prop-index="activeIndex"></NavMenu>
    <Title :prop-title="title"></Title>
    <Board :board-name="boardName"></Board>
  </div>
</template>

<script>
export default {
  name: 'General',
  data: function () {
    return {
      activeIndex: '2-1',
      title: '자유 게시판',
      boardName: 'general'
    }
  }
}
</script>

<style scoped>

</style>
