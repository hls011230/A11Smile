// pages/gainer_my/gainer_my.js

let app = getApp()

Component({
    data: {
        ename: '',
        info: '',
        protrait: '',
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
            this.getMess()
            this.getPortrait()
        }
    },
    methods: {
        getMess: function () {
            let _this = this
            wx.cloud.callContainer({
                "config": {
                    "env": "prod-9gy59jvo10e0946b"
                },
                "path": "/gainer/gainerAuthenticationSee",
                "header": {
                    "X-WX-SERVICE": "test-allsmile",
                    "content-type": "application/json",
                    "gid": app.globalData.gid
                },
                "method": "POST",
                "data": "",
                success: function (res) {
                    let data = res.data.data
                    _this.setData({
                        ename: data.enterprise_name,
                        info: data.resume
                    })
                }
            })
        },

        getPortrait:function(){
            let _this = this
            wx.cloud.callContainer({
                "config": {
                  "env": "prod-9gy59jvo10e0946b"
                },
                "path": "/gainer/showGainerIcon",
                "header": {
                  "X-WX-SERVICE": "test-allsmile",
                  "content-type": "application/json",
                  "gid": app.globalData.gid
                },
                "method": "POST",
                "data": "",
                success:function(res){
                    if (res.data.code == 0) {
                        _this.setData({
                            protrait: res.data.data
                        })
                    }else{
                        wx.showToast({
                          title: '服务器错误',
                          icon: 'error',
                          duration: '2000'
                        })
                    }
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
                    selected: 1
                })
            }
            this.getMess()
            this.getPortrait()
        }
    }
})