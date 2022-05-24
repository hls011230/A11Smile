// pages/user_certificate/user_certificate.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        serial: '',
        address: '',
        block_num: '',
        block_time: '',
        medical_examination_report_num: '',
        medical_history_num: '',
        birthday: '',
        gender: '',
        name: '',
        hash: ''
    },

    getData: function () {
        let _this = this
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/showDetailsCertificate",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "uid": app.globalData.uid
            },
            "method": "POST",
            "data": {
                "serial": _this.data.serial
            },
            success: function (res) {
                let data = res.data.data
                _this.setData({
                    address: data.address,
                    block_num: data.block_num,
                    block_time: data.block_time,
                    medical_examination_report_num: data.medical_examination_report_num,
                    medical_history_num: data.medical_history_num,
                    hash: data.serial
                })
            }
        })

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
                    birthday: data.birthday,
                    gender: data.gender
                })
            }
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let _this = this
        _this.setData({
            serial: options.serial
        })

        if (_this.data.serial != '') {
            _this.getData()
        }
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
        let _this = this
        if (_this.data.serial != '') {
            _this.getData()
        }
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