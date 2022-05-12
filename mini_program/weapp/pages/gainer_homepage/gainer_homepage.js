

Component({
    data: {
        data: '',
        index: [{
            pagePath: "/pages/gainer_homepage/gainer_homepage",
            iconPath: "/pages/images/202203141622511.png",
            selectedIconPath: "/pages/images/20220314162251.png",
            text: "首页"
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
                "path": "/gainer/gainerDisplayHomepage",
                "header": {
                    "X-WX-SERVICE": "test-allsmile",
                    "content-type": "application/json"
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
                    list: this.data.index,
                    selected: 0
                })
            }
            this.getData()
        }
    }
})
