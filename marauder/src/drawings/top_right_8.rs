use crate::proto::whiteboard::{drawing::Color, Drawing};

pub fn f(c: Color) -> Vec<Drawing> {
    vec![Drawing {
        xs: vec![
            1233.2878, 1192.7527, 1193.7347, 1194.5383, 1195.4312, 1196.5026, 1196.3241, 1196.2347,
            1196.3241, 1196.6812, 1197.3062, 1198.2883, 1199.7168, 1201.5918, 1203.7346, 1206.2346,
            1208.9132, 1211.681, 1214.2703, 1216.6809, 1218.6451, 1220.2523, 1221.5023, 1222.3059,
            1222.7523, 1222.9308, 1222.7523, 1222.038, 1220.788, 1219.0023, 1216.8595, 1214.1809,
            1211.2346, 1208.2882, 1205.5204, 1203.1096, 1201.2347, 1199.7168, 1198.6454, 1197.9312,
            1197.6633, 1197.7526, 1198.2883, 1199.1812, 1200.4312, 1201.949, 1203.556, 1205.431,
            1207.3954, 1209.181, 1211.056, 1212.3953, 1213.3774, 1214.0917, 1214.1809, 1214.0024,
            1213.1989, 1211.681, 1209.7167, 1207.0382, 1204.181, 1201.0562, 1198.199, 1195.6097,
            1193.5563, 1192.0383, 1191.0563, 1190.6099, 1190.6099, 1191.3241, 1192.5741, 1194.4491,
            1197.1276, 1200.2526, 1203.7346, 1207.3954, 1210.6989, 1213.6453,
        ],
        ys: vec![
            47.052227, 31.60624, 31.159822, 30.445559, 29.552727, 28.749178, 27.94563, 27.231363,
            26.427816, 25.624268, 24.731436, 23.660038, 22.410074, 21.160109, 19.999428, 18.749464,
            17.76735, 17.053083, 16.517385, 16.249535, 16.160252, 16.249535, 16.517385, 16.874517,
            17.410217, 18.124481, 19.195879, 20.62441, 22.677923, 25.17785, 28.213478, 31.516956,
            34.731148, 38.12391, 41.24882, 44.463013, 47.498642, 50.177135, 52.498497, 54.55201,
            56.426956, 58.301903, 59.909, 61.426815, 62.67678, 63.480328, 64.10531, 64.37316,
            64.46244, 64.55173, 64.46244, 64.194595, 63.926743, 63.56961, 63.123196, 62.408928,
            61.60538, 60.4447, 58.926888, 57.498356, 55.712692, 54.194878, 52.677067, 51.248535,
            49.820004, 48.03434, 45.980827, 43.659466, 41.15954, 38.570324, 36.070396, 33.213337,
            30.26699, 27.409931, 24.55287, 22.32079, 20.62441, 19.553013,
        ],
        pressures: vec![
            0 / 2,
            2882 / 2,
            3077 / 2,
            3275 / 2,
            3515 / 2,
            3678 / 2,
            3775 / 2,
            3841 / 2,
            3874 / 2,
            3882 / 2,
            3886 / 2,
            3889 / 2,
            3902 / 2,
            3828 / 2,
            3748 / 2,
            3754 / 2,
            3759 / 2,
            3763 / 2,
            3760 / 2,
            3752 / 2,
            3750 / 2,
            3748 / 2,
            3744 / 2,
            3747 / 2,
            3763 / 2,
            3800 / 2,
            3836 / 2,
            3856 / 2,
            3954 / 2,
            4044 / 2,
            4053 / 2,
            4057 / 2,
            4052 / 2,
            4052 / 2,
            4055 / 2,
            4055 / 2,
            4054 / 2,
            4051 / 2,
            4044 / 2,
            4032 / 2,
            4026 / 2,
            4027 / 2,
            4021 / 2,
            4007 / 2,
            4001 / 2,
            3997 / 2,
            4003 / 2,
            4009 / 2,
            4003 / 2,
            4010 / 2,
            4020 / 2,
            4039 / 2,
            4068 / 2,
            4095 / 2,
            4095 / 2,
            4101 / 2,
            4101 / 2,
            4101 / 2,
            4101 / 2,
            4101 / 2,
            4106 / 2,
            4106 / 2,
            4106 / 2,
            4112 / 2,
            4112 / 2,
            4112 / 2,
            4112 / 2,
            4116 / 2,
            4116 / 2,
            4116 / 2,
            4119 / 2,
            4121 / 2,
            4121 / 2,
            4122 / 2,
            4047 / 2,
            3578 / 2,
            2560 / 2,
            418 / 2,
        ],
        widths: vec![
            2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
            2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
            2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
        ],
        color: c as i32,
    }]
}
