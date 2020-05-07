<template>
  <el-table
    :data="this.$store.getters.getPosts"
    style="width: 100%; margin-top: 30px;"
    @cell-click="handlerClick"
    empty-text="게시글이 없습니다.">
    <el-table-column
      prop="date"
      label="Date"
      width="180">
    </el-table-column>
    <el-table-column
      prop="name"
      label="Name"
      width="180">
    </el-table-column>
    <el-table-column
      prop="title"
      label="Title">
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'Board',
  props: {
    propBoardName: String
  },
  data: function () {
    return {
      boardName: this.propBoardName
    }
  },
  created () {
    this.getAllPosts()
  },
  methods: {
    handlerClick: function (row) {
      this.$router.push(`/${this.boardName}/${row.id}`)
    },
    getAllPosts () {
      this.$store.dispatch('loadPostsFromServer', this.boardName)
    }
  }
}
</script>

<style>

  el-table th, .el-table tr, .el-table__header th {
    background-color: #34495e;
    color: #ecf0f1;
    cursor: pointer;
  }
  .el-table--enable-row-hover .el-table__body tr:hover>td{
    background-color: #7f8c8d !important;
  }

  .el-table__body tr:hover>td{
    background-color: #7f8c8d!important;
  }

  .el-table__body tr.current-row>td{
    background-color: #7f8c8d!important;
  }

  .el-table__empty-block {
    background-color: #34495e;
  }

  .el-table__empty-text {
    color: white;
  }
</style>
