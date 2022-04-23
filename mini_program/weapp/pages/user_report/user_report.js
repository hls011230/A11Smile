// pages/user_report/user_report.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        image: [],
        reports: [],
        temp: []
    },

    getData() {
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

    uploadReport: function () {
        var _this = this
        //让用户选择或拍摄一张照片
        wx.chooseImage({
            count: 1,
            sizeType: ['original', 'compressed'],
            sourceType: ['album', 'camera'],
            success(res) {
                //选择完成会先返回一个临时地址保存备用
                const tempFilePaths = res.tempFilePaths
                _this.setData({
                    image: tempFilePaths
                })
                wx.uploadFile({
                    filePath: _this.data.image[0],
                    name: 'uploadFile',
                    url: 'https://test-allsmile-1687181-1310014865.ap-shanghai.run.tcloudbase.com/user/uploadMedicalExaminationReport',
                    header: {
                        'content-type': 'application/x-www-form-urlencoded; charset=UTF-8',
                        'uid': app.globalData.uid
                    },
                    method: 'POST',
                    // formData: {
                    //     'user': 'test'
                    // },
                    success: function (res) {
                        if (res.data == '{"code":0,"data":"上传体检报告成功"}') {
                            wx.showToast({
                                title: '添加成功',
                                icon: 'success',
                                duration: 1500,
                                success: function () {
                                    let str = _this.data.image[0]
                                    let tem = str.slice(11)
                                    if (_this.data.reports == null) {
                                        _this.data.temp.unshift(tem)
                                        _this.setData({
                                            reports: _this.data.temp
                                        })
                                    } else {
                                        _this.data.reports.unshift(tem)
                                        let newReports = _this.data.reports
                                        _this.setData({
                                            reports: newReports
                                        })
                                    }
                                }
                            })
                        } else {
                            wx.showToast({
                                title: '添加失败',
                                icon: 'error',
                                duration: 1500,
                            })
                        }
                    },
                    fail: function (res) {
                        wx.showToast({
                            title: '添加失败',
                            icon: 'error',
                            duration: 1500,
                        })
                    }
                })
            },
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