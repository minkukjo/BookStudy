<template>
  <div class="block">
    <el-carousel :interval="4000" type="card" height="200px" style="margin-top: 200px">
      <el-carousel-item v-for="item in items" :key="item">
        <h3 class="medium">{{ item }}</h3>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      items: ['이상향에 도달하는 것에 실패함으로써 우리는 결국 스스로 누구인지 정의하게 되고 그 실패가 우리를 더욱 특별하게 만든다.', '어제와 똑같이 살면서 미래가 바뀌길 바라는 것은 정신병 초기 증세다.', 'Work hard, be kind and amazing things will happen', '인내력은 단순히 힘든 일을 참아낼 수 있는 능력이 아니라 그것을 영광으로 돌려놓을 수 있는 능력이다.']
    }
  },
  name: 'Carousel'
}
</script>

<style>
  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
    text-align: center;
  }

  .el-carousel__item:nth-child(2n) {
    background-color: #ffeaa7;
  }

  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
</style>
