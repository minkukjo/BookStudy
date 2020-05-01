<template>
  <el-table
    :data="tableData"
    style="width: 100%; margin-top: 30px;"
    @cell-click="handlerClick">
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
    boardName: String
  },
  data: function () {
    return {
      tableData: [{
        date: '2016-05-03',
        name: 'Harry',
        title: '즐거운 밤이다~',
        text: '나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~나아아아는야~~'
      }, {
        date: '2016-05-02',
        name: 'Harry',
        title: '코딩을 하면서 콜라를 마시고 밴드 음악을 든는건 정말 최고야',
        text: '매일매일 공부하는 개발자라네~~'
      }, {
        date: '2016-05-04',
        name: 'Harry',
        title: '이제 인턴이 한달 밖에 남지않았네',
        text: '힘들구먼~~~'
      }, {
        date: '2016-05-01',
        name: 'Harry',
        title: '휴 이거 언제 배포하냐 ㅋㅋ',
        text: '쥐엔장~~'
      }]
    }
  },
  methods: {
    handlerClick: function (row) {
      this.$router.push(`/main/${this.boardName}/${this.tableData.indexOf(row) + 1}`)
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
</style>
