// pages/gainer_check/user_check.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {

    },

    examine: function (e) {
        let _this = this
        let flag = e.currentTarget.dataset.flag
        let data = this.data
        if (flag == 'true') {
            wx.showModal({
                editable: true,
                placeholderText: '请输入要转账的金额',
                showCancel: true,
                title: '给分享者奖励',
                success: function (res) {
                    if (res.confirm == true && res.content != '') {
                        let ercnum = parseInt(res.content)
                        wx.cloud.callContainer({
                            "config": {
                                "env": "prod-9gy59jvo10e0946b"
                            },
                            "path": "/gainer/Examine",
                            "header": {
                                "X-WX-SERVICE": "test-allsmile",
                                "content-type": "application/json",
                                "gid": app.globalData.gid
                            },
                            "method": "POST",
                            "data": {
                                "certificate": data.certificate,
                                "medical_name": data.medical_name,
                                "whether": true,
                                "address": data.address,
                                "ercnum": ercnum
                            },
                            success: function (res) {
                                console.log(res)
                                wx.showToast({
                                    title: '支付奖励成功',
                                    icon: 'success',
                                    duration: 2000
                                })
                            }
                        })
                    }
                }
            })
        } else {
            wx.showModal({
                showCancel: true,
                title: '确认审核不通过',
                success: function (res) {
                    if (res.confirm == true) {
                        wx.cloud.callContainer({
                            "config": {
                                "env": "prod-9gy59jvo10e0946b"
                            },
                            "path": "/gainer/Examine",
                            "header": {
                                "X-WX-SERVICE": "test-allsmile",
                                "content-type": "application/json",
                                "gid": app.globalData.gid
                            },
                            "method": "POST",
                            "data": {
                                "certificate": data.certificate,
                                "medical_name": data.medical_name,
                                "whether": false,
                                "address": data.address,
                                "ercnum": 0
                            },
                            success: function (res) {
                                wx.showToast({
                                    title: '审核不通过',
                                    icon: 'error',
                                    duration: 2000
                                })
                            }
                        })
                    }
                }
            })
        }

    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let _this = this
        _this.setData({
            certificate: options.certificate
        })

        _this.setData({
            medical_name: options.medical_name
        })

        _this.setData({
            address: options.address
        })

    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide() {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload() {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {

    }
})