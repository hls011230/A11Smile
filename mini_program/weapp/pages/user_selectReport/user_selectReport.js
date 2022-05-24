// pages/user_selectReport/user_selectReport.js

let app = getApp()
let r = []
let m = []

Page({

    /**
     * 页面的初始数据
     */
    data: {
        reports: [],
        medicalHistory: [],
        array_medical_examination_report: [],
        array_medical_history: []
    },

    getReports() {
        let _this = this
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/viewAllMedicalExaminationReport",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "uid": app.globalData.uid
            },
            "method": "POST",
            "data": "",
            success: function (res) {
                _this.setData({
                    reports: res.data.data
                })
            }
        })
    },

    previewReports(e) {
        let _this = this
        let fileName = e.currentTarget.dataset.filename
        let value = e.detail.value
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/previewMedicalExaminationReport",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "uid": app.globalData.uid
            },
            "method": "POST",
            "data": {
                "file_name": fileName
            },
            success: function (res) {
                let data = res.data.data
                if (value.length != 0) {
                    r.push(data)
                    _this.setData({
                        array_medical_examination_report: r
                    })
                } else {
                    r.splice(r.indexOf(data), 1)
                    _this.setData({
                        array_medical_examination_report: r
                    })
                }
            }
        })
    },

    getMedical() {
        let _this = this
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/viewAllMedicalHistory",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "uid": app.globalData.uid
            },
            "method": "POST",
            "data": "",
            success: function (res) {
                _this.setData({
                    medicalHistory: res.data.data
                })
            }
        })
    },

    previewMedical(e) {
        let _this = this
        let fileName = e.currentTarget.dataset.filename
        let value = e.detail.value
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/previewMedicalHistory",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "uid": app.globalData.uid
            },
            "method": "POST",
            "data": {
                "file_name": fileName
            },
            success: function (res) {
                let data = res.data.data
                if (value.length != 0) {
                    m.push(data)
                    _this.setData({
                        array_medical_history: m
                    })
                } else {
                    m.splice(m.indexOf(data), 1)
                    _this.setData({
                        array_medical_history: m
                    })
                }
            }
        })
    },

    createCertificate: function () {
        let _this = this
        let medical = _this.data.array_medical_history
        let report = _this.data.array_medical_examination_report
        wx.showModal({
            editable: true,
            placeholderText: '请输入证书名称',
            showCancel: true,
            title: '请给需要生成的证书命名',
            success: function (res) {
                if (res.confirm == true && res.content != '') {
                    let fileName = res.content
                    if (report == [] && medical == []) {
                        wx.showToast({
                            title: '请至少选择一项',
                            icon: 'error',
                            duration: 2000
                        })
                    } else {
                        wx.cloud.callContainer({
                            "config": {
                                "env": "prod-9gy59jvo10e0946b"
                            },
                            "path": "/user/createCertificate",
                            "header": {
                                "X-WX-SERVICE": "test-allsmile",
                                "content-type": "application/json",
                                "uid": app.globalData.uid
                            },
                            "method": "POST",
                            "data": {
                                "array_medical_history": medical,
                                "array_medical_examination_report": report,
                                "certificate_name": fileName
                            },
                            success: function (res) {
                                if (res.data.data == '证书生成成功') {
                                    wx.showToast({
                                        title: '正在生成证书',
                                        icon: 'loading',
                                        duration: 8000,
                                        success: function () {
                                            setTimeout(() => {
                                                wx.showToast({
                                                    title: '证书生成成功',
                                                    icon: 'success',
                                                    duration: 2000,
                                                    success: function () {
                                                        setTimeout(() => {
                                                            wx.switchTab({
                                                                url: '/pages/user_data/user_data',
                                                            })
                                                        }, 2000);
                                                    }
                                                })
                                            }, 8000);
                                        }
                                    })
                                } else {
                                    wx.showToast({
                                        title: '证书生成失败',
                                        icon: 'error',
                                        duration: 2000,
                                        success: function () {
                                            wx.switchTab({
                                                url: '/pages/user_data/user_data',
                                            })
                                        }
                                    })
                                }
                            }
                        })
                    }
                }
            }
        })


    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        this.getReports()
        this.getMedical()
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
        this.getReports()
        this.getMedical()
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