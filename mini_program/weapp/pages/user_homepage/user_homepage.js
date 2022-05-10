
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
        pageData: '',
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

        goDetailsPage: function (e) {
            let _this = this
            _this.setData({
                medical_name: e.currentTarget.dataset.medical_name,
                address: e.currentTarget.dataset.address
            })
            wx.navigateTo({
                url: '/pages/user_uploading/user_uploading',
                success: function (res) {
                    res.eventChannel.emit('medical_name', { data: _this.data.medical_name})
                    res.eventChannel.emit('address', { data: _this.data.address})
                }
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