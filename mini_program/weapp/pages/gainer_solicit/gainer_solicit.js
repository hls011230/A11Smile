// pages/gainer_solicit/gainer_solicit.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        min: '',
        max: '',
        account: '',
        medical_name: '',
        medical_need: '',
        requirement_description: ''
    },

    setMin: function (e) {
        this.setData({
            min: e.detail.value
        })
    },

    setMax: function (e) {
        this.setData({
            max: e.detail.value
        })
    },

    setAccount: function (e) {
        this.setData({
            account: e.detail.value
        })
    },

    setName: function (e) {
        this.setData({
            medical_name: e.detail.value
        })
    },

    setNeed: function (e) {
        this.setData({
            medical_need: e.detail.value
        })
    },

    setRequire: function (e) {
        this.setData({
            requirement_description: e.detail.value
        })
    },

    deal: function (e) {
        wx.navigateTo({
            url: '/pages/user_deal/user_deal',
        })
    },

    solicit: function () {
        let _this = this
        let min = parseInt(_this.data.min)
        let max = parseInt(_this.data.max)
        let account = parseInt(_this.data.account)
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/gainer/ReleaseMedicalInformation",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "gid": app.globalData.gid
            },
            "method": "POST",
            "data": {
                "min": min,
                "max": max,
                "account": account,
                "medical_name": _this.data.medical_name,
                "medical_need": _this.data.medical_need,
                "requirement_description": _this.data.requirement_description
            },
            success: function (res) {
                if (res.data.data == '征求者发布医疗信息成功') {
                    wx.showToast({
                        title: '发布成功',
                        icon: 'success',
                        duration: 2000,
                        success: function () {
                            setTimeout(() => {
                                wx.switchTab({
                                    url: '/pages/gainer_homepage/gainer_homepage',
                                })
                            }, 2000);
                        }
                    })
                } else {
                    wx.showToast({
                        title: '发布失败',
                        icon: 'error',
                        duration: 2000
                    })
                }
            }
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {

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