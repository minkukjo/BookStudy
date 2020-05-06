<template>
  <div class="study">
    <NavMenu :prop-index="activeIndex"></NavMenu>
    <Title :prop-title="title"></Title>
    <Board :prop-board-name="boardName"></Board>
  </div>
</template>

<script>
export default {
  name: 'Study',
  data: function () {
    return {
      activeIndex: '2-2',
      title: '스터디 모집',
      boardName: 'study'
    }
  }
}
</script>

<style scoped>

</style>
