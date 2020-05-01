<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">
      <el-row>
        <el-col :span="18">
          <div class="title"><i v-on:click="goBack" class="el-icon-back"></i> {{tableData.title}}</div>
        </el-col>
        <el-col :span="6"><div class="writerDate">{{tableData.name}} </div><div class="writerDate">{{tableData.date}}</div></el-col>
      </el-row>
    </div>
    <div class="textBox">
      <div class="text item">
        {{tableData.text}}
      </div>
    </div>
  </el-card>
</template>

<script>
import Board from './Board'
export default {
  name: 'Detail',
  props: {
    propBoardName: String
  },
  data () {
    const index = this.$route.params.id - 1
    return {
      tableData: Board.data().tableData[index],
      boardName: this.propBoardName
    }
  },
  methods: {
    goBack () {
      if (this.boardName !== undefined) {
        this.$router.push(`/main/${this.boardName}`)
      }
    }
  }
}
</script>

<style scoped>
  .el-icon-back {
    cursor: pointer;
  }
  .text {
    font-size: 20px;
    text-align: left;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix {
    text-align: left;
    color: white;
  }

  .box-card {
    margin-top: 30px;
    width: 1300px;
    height: 700px;
  }

  .title {
    margin-top: 30px;
    font-size: 30px;
    font-weight: bold;
    text-align: left;
  }

  .writerDate {
    margin-top: 20px;
    margin-bottom: 20px;
    font-size: 20px;
    text-align: right;
  }
</style>
