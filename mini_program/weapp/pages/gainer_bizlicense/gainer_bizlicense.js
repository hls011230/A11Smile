// pages/gainer_bizlicense/gainer_bizlicense.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        image: []
    },

    upload: function () {
        //让用户选择或拍摄一张照片
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
            },
        })
    },

    verifyBizlicense: function () {
        let _this = this
        let image = _this.data.image
        if (image[0] == undefined) {
            wx.showToast({
                title: '请上传营业执照',
                icon: 'error',
                duration: 1000,
            })
        }else{
            wx.uploadFile({
                filePath: this.data.image[0],
                name: 'uploadBizlicense',
                url: 'https://test-allsmile-1687181-1310014865.ap-shanghai.run.tcloudbase.com/gainer/register/verifyBizlicense',
                header: {
                    'content-type': 'application/x-www-form-urlencoded; charset=UTF-8',
                },
                method: 'POST',
                success: function (res) {
                    let data = JSON.parse(res.data)
                    if(data.code == 0){
                        app.globalData.gid = data.data
                        wx.showToast({
                          title: '验证成功',
                          icon: 'success',
                          duration: 1000,
                          success: function(){
                              wx.navigateTo({
                                url: '/pages/gainer_register/gainer_register',
                              })
                          }
                        })
                    }else{
                        wx.showToast({
                          title: '验证失败',
                          icon: 'error',
                          duration: 1000
                        })
                    }
                }
            })
        }
    },

    previewImage(e) {
        let current = e.target.dataset.src
        wx.previewImage({
            current: current, // 当前显示图片的http链接
            urls: this.data.image // 需要预览的图片http链接列表
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {

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