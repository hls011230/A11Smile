// pages/user_autonym/user_autonym.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        image: []
    },

    upload_picture: function (name) {
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
                console.log(tempFilePaths)
                //将照片上传至云端需要刚才存储的临时地址
                // wx.cloud.uploadFile({
                //     cloudPath: 'images/test/' + new Date().getTime() + "_" + Math.floor(Math.random() * 1000) + '.jpg',
                //     filePath: tempFilePaths[0],
                //     success(res) {
                //         console.log(res.fileID)
                //         console.log('上传图片成功')
                //         wx.showToast({
                //             title: '上传成功',
                //             icon: 'success',
                //             duration: 1000,
                //             success: function () {
                //                 setTimeout(
                //                     function () {
                //                         wx.switchTab({
                //                             url: '/pages/user_data/user_data',
                //                         })
                //                     }, 1000)
                //             }
                //         })
                //     },
                //     fail(res) {
                //         wx.showToast({
                //             title: '上传失败',
                //             icon: 'fail',
                //             duration: 1000
                //         })
                //     }
                // })
            }
        })
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