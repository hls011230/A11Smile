// pages/user_balance/user_balance.js
let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        eth_balance: '',
        as_balance: '',
        token: 'AS',
    },

    //切换显示代币余额
    switchToken(e) {
        if (this.data.token == 'AS') {
            this.setData({
                token: 'ETH'
            })
        } else {
            this.setData({
                token: 'AS'
            })
        }
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        let _this = this
        //获取AS余额
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/gainer/CheckTheAS",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "gid": app.globalData.gid
            },
            "method": "POST",
            "data": "",
            success: function (res) {
                _this.setData({
                    as_balance: res.data.data
                })
            }
        })

        //获取ETH余额
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/gainer/CheckTheBalance",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "gid": app.globalData.gid
            },
            "method": "POST",
            "data": "",
            success: function (res) {
                _this.setData({
                    eth_balance: res.data.data
                })
            }
        })
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