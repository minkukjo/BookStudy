<template>
  <div class="write" style="margin-top: 30px">
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="120px" class="demo-ruleForm" style="width: 40%" >
      <el-form-item label="제목" prop="title">
        <el-input type="title" v-model="ruleForm.title" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="내용" prop="text">
        <el-input type="textarea" v-model="ruleForm.text"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">Submit</el-button>
        <el-button @click="resetForm('ruleForm')">Reset</el-button>
      </el-form-item>
    </el-form>
  </div>

</template>

<script>
export default {
  name: 'write',
  data () {
    return {
      ruleForm: {
        title: '',
        text: ''
      },
      rules: {
        title: [
          { required: true, message: '제목을 입력해주세요.', trigger: 'blur' }
        ],
        text: [
          { required: true, message: '내용을 입력해주세요.', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log(this.ruleForm)
          this.$http.post(this.$http_url + '/api/write', this.ruleForm)
            .then((response) => {
              this.$router.push('/main/general')
            })
          alert('작성 완료!')
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
    }
  }

}
</script>

<style>
  .el-form-item label {
    color:white;
  }

  .el-textarea textarea {
    height: 300px;
  }
</style>
