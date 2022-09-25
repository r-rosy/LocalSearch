<template>
  <div id="view">
    <div id="new">
      <div class="text">新增索引键值对</div>
    <el-input v-model="inputkey1" placeholder="请输入新增key"></el-input>
    <el-input v-model="inputvalue1" placeholder="请输入新增value"></el-input>
    <el-button type="primary" @click="New">新增</el-button>
    
    </div>
    
    <div class="text">搜索索引键值对</div>
    <div id="find">
      <el-table
    :data="Data"
    stripe
    style="width: 100%">
    <el-table-column
      prop="id"
      label="编号"
      width="180">
    </el-table-column>
    <el-table-column
      prop="host"
      label="索引"
      width="180">
    </el-table-column>
    <el-table-column
      prop="guest"
      label="值"
      width="180">
    </el-table-column>
  </el-table>
  <el-input v-model="query" placeholder="请输入索引"></el-input>
    <el-button type="primary" @click="SearchHost">搜索</el-button>
    </div>
    <div class="text">删除host</div>
    <el-input v-model="deletehost" placeholder="请输入删除的host"></el-input>
    <el-button type="primary" @click="DeleteHost">删除</el-button>

    <div class="text">删除host及guest</div>
    <el-input v-model="GuestId" placeholder="请输入删除guest的ID"></el-input>
    <el-button type="primary" @click="DeleteGuest">删除</el-button>

    
    </div>

    
</template>

<script>
import axios from 'axios'
export default {
  name: 'HelloWorld',
  data () {
    return {
      Data: [],
      inputkey1: '',
      inputvalue1: '',
      deletehost: '',
      query: '',
      GuestId: '',
      }
  }
  ,
  methods: {
      New: function() {
        var pointer=this
        if(this.inputkey1==''||this.inputvalue1=='')
        {
            this.$alert("请同时输入完整的key-value")
            return
        }
            axios.post('/api/v1/map?host='+pointer.inputkey1+'&guest='+pointer.inputvalue1).then(function (response) {
                console.log(response)
                    if(response.data.code==200)
                    {
                      pointer.$message("添加成功")
                      pointer.$data =  pointer.$options.data()
                      Object.assign(pointer.$data, pointer.$options.data())
                    }else {
                      pointer.$alert("添加失败")
                      pointer.$data =  pointer.$options.data()
                      Object.assign(pointer.$data, pointer.$options.data())
                    }
                })
            .catch((error) =>{
    console.log(error);
  });
      },
      SearchHost: function() {
        var pointer=this
        if(this.query=='')
        {
            this.$alert("请输入完整的索引")
            return
        }
            axios.get('/api/v1/map?host='+pointer.query).then(function (response) {
                console.log(response)
                    if(response.data.code==200)
                    {
                      var res=response.data.data
                      if(res==null)
                      {
                        pointer.$alert("查询为空")
                        return
                      }
                      for (var i=0;i<res.length;i++)
                      {
                        pointer.Data.push({id:res[i].id,guest:res[i].data,host:pointer.query})
                      }
                      pointer.$forceUpdate()
                    }else {
                      pointer.$alert("出错失败")
                    }
                })
            .catch((error) => {
            console.log(error);
            
  });
      },
      DeleteGuest: function() {
        var pointer=this
        
        axios.delete('/api/v1/map',{params:{value:pointer.GuestId,method:"guest"}}).then(function (response) {
            console.log(response)
                if(response.data.code==200)
                {
                  pointer.$message("删除成功")
                  pointer.$data =  pointer.$options.data()
                  Object.assign(pointer.$data, pointer.$options.data())
                }else {
                  pointer.$alert("删除失败")
                }
                pointer.$forceUpdate()
            })
        .catch(function (error) {
console.log(error);
});
      },
      DeleteHost: function() {
        var pointer=this
        if(this.deletehost=='')
        {
            this.$alert("请输入完整的索引")
            return
        }
        axios.delete('/api/v1/map?value='+pointer.deletehost+'&method=host').then(function (response) {
            console.log(response)
                if(response.data.code==200)
                {
                  pointer.$message("删除成功")
                  pointer.$data =  pointer.$options.data()
                  Object.assign(pointer.$data, pointer.$options.data())
                }else {
                  pointer.$alert("删除失败")
                }
                pointer.$forceUpdate()
            })
        .catch((error) =>{
console.log(error);
});
  },
}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}

</style>
