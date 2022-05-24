// pages/user_selectCertificate/user_selectCertificate.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        certificatelist: '',
        medical_name: '',
        soliciter: '',
        certificate: '',
        name: ''
    },

    getData: function () {
        let _this = this
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/showAllCertificate",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "uid": app.globalData.uid
            },
            "method": "POST",
            "data": {},
            success: function (res) {
                _this.setData({
                    certificatelist: res.data.data
                })
            }
        })
    },

    selectCertificate: function (e) {
        let _this = this
        let value = e.detail.value
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
                "serial": value
            },
            success: function (res) {
                console.log(res)
                let data = res.data.data
                _this.setData({
                    certificate: data.serial,
                    name: value
                })
            }
        })
    },

    submitCertificate: function () {
        let _this = this
        let certificate = _this.data.certificate
        let soliciter = _this.data.soliciter
        let medical_name = _this.data.medical_name
        if (certificate == '') {
            wx.showToast({
                title: '请选择要投递的证书',
                icon: 'error',
                duration: 2000
            })
        } else {
            wx.cloud.callContainer({
                "config": {
                    "env": "prod-9gy59jvo10e0946b"
                },
                "path": "/user/SubmitCertificate",
                "header": {
                    "X-WX-SERVICE": "test-allsmile",
                    "content-type": "application/json",
                    "uid": app.globalData.uid
                },
                "method": "POST",
                "data": {
                    "certificate_": certificate,
                    "soliciter_": soliciter,
                    "medical_name_": medical_name,
                    "department": "其他"
                },
                success: function (res) {
                    console.log(res)
                    if (res.data.data == '证书上传成功') {
                        wx.showToast({
                            title: '投递证书成功',
                            icon: 'success',
                            duration: 2000,
                            success: function () {
                                setTimeout(() => {
                                    wx.switchTab({
                                        url: '/pages/user_homepage/user_homepage',
                                    })
                                }, 2000);

                            }
                        })
                    } else {
                        wx.showToast({
                            title: '投递证书失败',
                            icon: 'error',
                            duration: 2000
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
        this.getData()
        this.setData({
            medical_name: options.medical_name,
            soliciter: options.address
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
        this.getData()
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