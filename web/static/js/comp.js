compMainHeader = Vue.extend({
    template: `
<el-header id="main-header">
    <p id="main-title"> BiliShow </p>
    <div id="main-box">
        <el-dropdown>
            <img id="main-face"
                src="https://gimg3.baidu.com/search/src=http%3A%2F%2Fgips0.baidu.com%2Fit%2Fu%3D3462928813%2C851336810%26fm%3D3030%26app%3D3030%26f%3DJPEG%3Fw%3D121%26h%3D74%26s%3D37F6EC36C494FF90237D87C403005025&refer=http%3A%2F%2Fwww.baidu.com&app=2021&size=f242,150&n=0&g=0n&q=100&fmt=auto?sec=1697562000&t=75c54148f2f3f403219190899b52edc1" />
            <el-dropdown-menu slot="dropdown">
               <el-dropdown-item class="main-dropdown-item" v-for="(item,index) in itemList"  :key="item.id"  @click.native="gotoLink(item)">{{item.text}}</el-dropdown-item>
            </el-dropdown-menu>
        </el-dropdown>
    </div>
    </el-header>
`,
    data(){
        return {
            itemList:[
            ]
        }
    },
    methods:{
        setItemList(){
            this.itemList = globalMainBoxItemList
        },
        gotoLink(item){
            location.href = item.link
        }
    },
    mounted(){
        this.setItemList()
    }
})
Vue.component("main-header", compMainHeader);

compMainAside = Vue.extend({
    template: `
<el-aside id="main-aside" width="200px">
    <el-menu :default-openeds="['1', '2']">
        <el-submenu index="1">
            <template slot="title"><i class="el-icon-message"></i>静态数据</template>
            <el-menu-item index="1-1">UP主数据</el-menu-item>
            <el-menu-item index="1-2">视频数据</el-menu-item>
        </el-submenu>
        <el-submenu index="2">
            <template slot="title"><i class="el-icon-menu"></i>动态数据</template>
            <el-menu-item index="2-1">UP主粉丝量任务</el-menu-item>
            <el-menu-item index="2-2">UP主粉丝量追踪</el-menu-item>
            </el-submenu>
        </el-submenu>
    </el-menu>
</el-aside>
`,

})
Vue.component("main-aside", compMainAside);