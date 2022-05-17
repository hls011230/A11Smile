let app = getApp()

Component({
    data: {
        data: '',
        pageData: [],
        iconUrl: '',
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
                "path": "/gainer/gainerDisplayHomepage",
                "header": {
                  "X-WX-SERVICE": "test-allsmile",
                  "content-type": "application/json",
                  "gid": app.globalData.gid
                },
                "method": "POST",
                "data": "",
                success:function(res){
                    let data = res.data.data
                    _this.setData({
                        pageData: data
                    })
                }
              })
              wx.cloud.callContainer({
                "config": {
                  "env": "prod-9gy59jvo10e0946b"
                },
                "path": "/gainer/showGainerIcon",
                "header": {
                  "X-WX-SERVICE": "test-allsmile",
                  "content-type": "application/json",
                  "gid": "1"
                },
                "method": "POST",
                "data": "",
                success:function(res){
                    _this.setData({
                        iconUrl: res.data.data
                    })
                }
              })
        },
        goDetailsPage: function (e) {
            let _this = this
            let medicalName = e.currentTarget.dataset.medicalname
            wx.navigateTo({
                url: '/pages/gainer_certificatelist/gainer_certificatelist?MedicalName='+medicalName,
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
