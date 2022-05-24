
let app = getApp()

Component({
    /**
     * 组件的属性列表
     */
    properties: {

    },

    lifetimes: {
        attached: function () {
            this.getData()
        }
    },

    /**
     * 组件的初始数据
     */
    data: {
        pageData:[],
        medical_name: '',
        address: '',
        home: [{
            pagePath: "/pages/user_homepage/user_homepage",
            text: "首页",
            iconPath: "/pages/images/202203141622511.png",
            selectedIconPath: "/pages/images/20220314162251.png"
        },
        {
            pagePath: "/pages/user_data/user_data",
            text: "数据管理",
            iconPath: "/pages/images/916920790148990582.png",
            selectedIconPath: "/pages/images/169545831883691004.png"
        },
        {
            pagePath: "/pages/user_my/user_my",
            text: "我的",
            iconPath: "/pages/images/202203141622512.png",
            selectedIconPath: "/pages/images/202203141622513.png"
        }]
    },

    /**
     * 组件的方法列表
     */
    methods: {
        getData: function () {
            let _this = this
            wx.cloud.callContainer({
                "config": {
                    "env": "prod-9gy59jvo10e0946b"
                },
                "path": "/user/userDisplayHomepage",
                "header": {
                    "X-WX-SERVICE": "test-allsmile",
                    "content-type": "application/json",
                },
                "method": "POST",
                "data": {},
                success: function (res) {
                    let data = res.data.data 
                    _this.setData({
                        pageData: data
                    })
                }
            })
        },

        classify: function(e){
            let _this = this
            let classify = e.currentTarget.dataset.classify
            wx.navigateTo({
              url: '/pages/user_classify/user_classify?classify='+classify,
            })
        },

        goDetailsPage: function (e) {
            let _this = this
            _this.setData({
                medical_name: e.currentTarget.dataset.medical_name,
                address: e.currentTarget.dataset.address,
                hospital_name: e.currentTarget.dataset.hospital_name
            })
            let medical_name = e.currentTarget.dataset.medical_name
            let address = e.currentTarget.dataset.address
            let hospital_name = e.currentTarget.dataset.hospital_name
            wx.navigateTo({
                url: "/pages/user_uploading/user_uploading?medical_name="+medical_name+"&address="+address+"&hospital_name="+hospital_name
            })
        }
    },
    pageLifetimes: {
        show() {
            if (typeof this.getTabBar === 'function' &&
                this.getTabBar()) {
                this.getTabBar().setData({
                    list: this.data.home,
                    selected: 0
                })
            }
            this.getData()
        }
    }
})