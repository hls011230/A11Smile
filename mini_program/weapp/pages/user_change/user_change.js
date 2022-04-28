// pages/user_change/user_change.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        uname: '',
        info: ''
    },

    getData: function () {
        let _this = this
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/userAuthenticationSee",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "uid": app.globalData.uid
            },
            "method": "POST",
            "data": "",
            success: function (res) {
                let data = res.data.data
                _this.setData({
                    uname: data.uname,
                    info: data.resume,
                })
            }
        })
    },

    changeName: function () {
        let _this = this
        wx.showModal({
            editable: true,
            placeholderText: '请输入新用户名',
            showCancel: true,
            title: '请设置新的用户名',
            success: (res) => {
                if (res.confirm == true && res.content != '') {
                    let newName = res.content
                    wx.cloud.callContainer({
                        "config": {
                          "env": "prod-9gy59jvo10e0946b"
                        },
                        "path": "/user/editUserName",
                        "header": {
                          "X-WX-SERVICE": "test-allsmile",
                          "content-type": "application/json",
                          "uid": app.globalData.uid
                        },
                        "method": "POST",
                        "data": {
                          "uname": newName
                        },success:(res) => {
                            if (res.data.data == '修改成功!') {
                                wx.showToast({
                                    title: '修改成功',
                                    icon: 'success',
                                    duration: 1500,
                                    success: (res) => {
                                        _this.setData({
                                            uname:newName
                                        })
                                    }
                                })
                            }else{
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

    changeInfo: function () {
        let _this = this
        wx.showModal({
            editable: true,
            placeholderText: '请输入新的个人简介',
            showCancel: true,
            title: '请设置新的个人简介',
            success: (res) => {
                if (res.confirm == true && res.content != '') {
                    let newInfo = res.content
                    wx.cloud.callContainer({
                        "config": {
                          "env": "prod-9gy59jvo10e0946b"
                        },
                        "path": "/user/editUserResume",
                        "header": {
                          "X-WX-SERVICE": "test-allsmile",
                          "content-type": "application/json",
                          "uid": app.globalData.uid
                        },
                        "method": "POST",
                        "data": {
                          "resume": newInfo
                        },success:(res) => {
                            if (res.data.data == '修改成功!') {
                                wx.showToast({
                                    title: '修改成功',
                                    icon: 'success',
                                    duration: 1500,
                                    success: (res) => {
                                        _this.setData({
                                            info:newInfo
                                        })
                                    }
                                })
                            }else{
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

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.getData()
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