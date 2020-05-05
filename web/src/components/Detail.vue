<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">
      <el-row>
        <el-col :span="18">
          <div class="title"><i v-on:click="goBack" class="el-icon-back"></i> {{postData.title}}</div>
        </el-col>
        <el-col :span="6"><div class="writerDate">{{postData.name}}</div><div class="writerDate">{{postData.date}}</div></el-col>
      </el-row>
    </div>
    <div class="textBox">
      <div class="text item">
        {{postData.text}}
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
    return {
      postData: null,
      boardName: this.propBoardName
    }
  },
  methods: {
    goBack () {
      if (this.boardName !== undefined) {
        this.$router.push(`/main/${this.boardName}`)
      }
    }
  },
  created () {
    const index = this.$route.params.id
    console.log(index)

    const findPost = Board.methods.findPost(index)
    console.log(findPost)
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
