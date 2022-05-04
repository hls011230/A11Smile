// pages/gainer_change/gainer_change.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        info: '',
        protrait: ''
    },

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
                    info: data.resume
                })
            }
        })
    },

    getPortrait: function () {
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
            success: function (res) {
                if (res.data.code == 0) {
                    _this.setData({
                        protrait: res.data.data
                    })
                } else {
                    wx.showToast({
                        title: '服务器错误',
                        icon: 'error',
                        duration: '2000'
                    })
                }
            }
        })
    },

    changeResume: function () {
        let _this = this
        wx.showModal({
            editable: true,
            placeholderText: '请输入新的简介',
            showCancel: true,
            title: '请设置新的简介',
            success: function (res) {
                if (res.confirm == true && res.content != '') {
                    let newResume = res.content
                    wx.cloud.callContainer({
                        "config": {
                            "env": "prod-9gy59jvo10e0946b"
                        },
                        "path": "/gainer/gainerEdit",
                        "header": {
                            "X-WX-SERVICE": "test-allsmile",
                            "content-type": "application/json",
                            "gid": app.globalData.gid
                        },
                        "method": "POST",
                        "data": {
                            "resume": newResume
                        },
                        success: function (res) {
                            if (res.data.data == '修改成功') {
                                wx.showToast({
                                    title: '修改成功',
                                    icon: 'success',
                                    duration: 1500,
                                    success: (res) => {
                                        _this.setData({
                                            info: newResume
                                        })
                                    }
                                })
                            } else {
                                wx.showToast({
                                    title: '修改失败',
                                    icon: 'error',
                                    duration: 1500,
                                })
                            }
                        }
                    })
                }
            }
        })
    },

    changeProtrait: function () {
        let _this = this
        wx.chooseImage({
            count: 1,
            sizeType: ['original', 'compressed'],
            sourceType: ['album', 'camera'],
            success: function (res) {
                const tempFilePaths = res.tempFilePaths
                _this.setData({
                    image: tempFilePaths
                }),
                    wx.uploadFile({
                        filePath: _this.data.image[0],
                        name: 'editImage',
                        url: 'https://test-allsmile-1687181-1310014865.ap-shanghai.run.tcloudbase.com/gainer/editGainerIcon',
                        header: {
                            'content-type': 'application/x-www-form-urlencoded; charset=UTF-8',
                            'gid': app.globalData.gid
                        },
                        method: 'POST',
                        success:function(res){
                            if (res.data.data == '修改成功') {
                                wx.showToast({
                                  title: '修改成功',
                                  icon: 'success',
                                  duration: 2000,
                                  success:function(res){
                                      _this.setData({
                                          protrait: _this.data.image[0]
                                      })
                                  }
                                })
                            }else{
                                wx.showToast({
                                  title: '修改失败',
                                  icon: 'error',
                                  duration: 2000
                                })
                            }
                        }
                    })
            }
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.getMess()
        this.getPortrait()
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {
        this.getMess()
        this.getPortrait()
    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide: function () {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload: function () {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh: function () {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom: function () {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {

    }
})