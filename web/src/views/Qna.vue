<template>
  <div class="qna">
    <Title :prop-title="title"></Title>
    <Board :prop-board-name="boardName"></Board>
  </div>
</template>

<script>
export default {
  name: 'Qna',
  data: function () {
    return {
      title: '질문 게시판',
      boardName: 'qna'
    }
  }
}
</script>

<style scoped>

</style>
