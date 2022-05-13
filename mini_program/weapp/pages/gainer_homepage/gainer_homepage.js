

Component({
    data: {
        data: '',
        index: [{
            pagePath: "/pages/gainer_homepage/gainer_homepage",
            iconPath: "/pages/images/202203141622511.png",
            selectedIconPath: "/pages/images/20220314162251.png",
            text: "发布"
        }, {
            pagePath: "/pages/gainer_my/gainer_my",
            iconPath: "/pages/images/202203141622512.png",
            selectedIconPath: "/pages/images/202203141622513.png",
            text: "我的"
        }]
    },

    lifetimes: {
        attached: function () {
            this.getData()
        }
    },

    methods: {
        getData: function () {
            let _this = this
            wx.cloud.callContainer({
                "config": {
                    "env": "prod-9gy59jvo10e0946b"
                },
                "path": "/gainer/ViewCertificate",
                "header": {
                    "X-WX-SERVICE": "test-allsmile",
                    "content-type": "application/json",
                    "gid": "1"
                },
                "method": "POST",
                "data": "",
                success: function (res) {
                    _this.setData({
                        data: res.data.data
                    })
                }
            })
        },
        goDetailsPage: function (e) {
            let _this = this
            wx.navigateTo({
                url: '/pages/gainer_check/user_check?',
            })
        }
    },
    pageLifetimes: {
        show() {
            if (typeof this.getTabBar === 'function' &&
                this.getTabBar()) {
                this.getTabBar().setData({
                    list: this.data.index,
                    selected: 0
                })
            }
            this.getData()
        }
    }
})
