<template>
  <div class="general">
    <Title :prop-title="title"></Title>
    <Board :prop-board-name="boardName"></Board>
  </div>
</template>

<script>
export default {
  name: 'General',
  data: function () {
    return {
      title: '자유 게시판',
      boardName: 'general'
    }
  }
}
</script>

<style scoped>

</style>
