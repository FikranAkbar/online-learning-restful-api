package seeder

import (
	"database/sql"
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterElearningModuleSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterElearningModuleSeeder(cfg gormseeder.SeederConfiguration) *MasterElearningModuleSeeder {
	return &MasterElearningModuleSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitElearningModuleData() []entity.MasterElearningModule {
	return []entity.MasterElearningModule{
		//#Module Intro to Entrepreneurship
		{
			ID:          1,
			CourseId:    1,
			ModuleTitle: "Title 1",
			ModuleOverview: sql.NullString{
				String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship.",
				Valid:  true,
			},
			ModuleDescription: sql.NullString{
				String: "<h2>JUDUL</h2>\n                    <p>Terdiri dari beberapa bagian part pembelajaran untuk mendapatkan pengetahuan dan pemahaman baru tentang bagaimana cara membangun usaha, hal-hal apa saja yang dibutuhkan dalam membangun usaha baik secara internal maupun eksternal,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Memulai dari awal dan dasar-dasar yang perlu dibangun atau fundamental dalam membangun Enterpreneurship dan juga melihat perkembangan zaman dan teknologi yang berdasarkan IoT sebagai pusat pembelajaran utama dalam masyarakat.</p>",
				Valid:  true,
			},
			IsPublished: true,
		},
		{
			ID:          2,
			CourseId:    1,
			ModuleTitle: "Title 2",
			ModuleOverview: sql.NullString{
				String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship.",
				Valid:  true,
			},
			ModuleDescription: sql.NullString{
				String: "<h2>JUDUL</h2>\n                    <p>Terdiri dari beberapa bagian part pembelajaran untuk mendapatkan pengetahuan dan pemahaman baru tentang bagaimana cara membangun usaha, hal-hal apa saja yang dibutuhkan dalam membangun usaha baik secara internal maupun eksternal,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Memulai dari awal dan dasar-dasar yang perlu dibangun atau fundamental dalam membangun Enterpreneurship dan juga melihat perkembangan zaman dan teknologi yang berdasarkan IoT sebagai pusat pembelajaran utama dalam masyarakat.</p>",
				Valid:  true,
			},
			IsPublished: true,
		},
		{
			ID:          3,
			CourseId:    1,
			ModuleTitle: "Title 3",
			ModuleOverview: sql.NullString{
				String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship.",
				Valid:  true,
			},
			ModuleDescription: sql.NullString{
				String: "<h2>JUDUL</h2>\n                    <p>Terdiri dari beberapa bagian part pembelajaran untuk mendapatkan pengetahuan dan pemahaman baru tentang bagaimana cara membangun usaha, hal-hal apa saja yang dibutuhkan dalam membangun usaha baik secara internal maupun eksternal,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Memulai dari awal dan dasar-dasar yang perlu dibangun atau fundamental dalam membangun Enterpreneurship dan juga melihat perkembangan zaman dan teknologi yang berdasarkan IoT sebagai pusat pembelajaran utama dalam masyarakat.</p>",
				Valid:  true},
			IsPublished: true,
		},
		{
			ID:          4,
			CourseId:    1,
			ModuleTitle: "Title 4",
			ModuleOverview: sql.NullString{
				String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship.",
				Valid:  true,
			},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Terdiri dari beberapa bagian part pembelajaran untuk mendapatkan pengetahuan dan pemahaman baru tentang bagaimana cara membangun usaha, hal-hal apa saja yang dibutuhkan dalam membangun usaha baik secara internal maupun eksternal,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Memulai dari awal dan dasar-dasar yang perlu dibangun atau fundamental dalam membangun Enterpreneurship dan juga melihat perkembangan zaman dan teknologi yang berdasarkan IoT sebagai pusat pembelajaran utama dalam masyarakat.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Intermediate to Entrepreneurship
		{
			ID:          5,
			CourseId:    5,
			ModuleTitle: "Title 1",
			ModuleOverview: sql.NullString{
				String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship.",
				Valid:  true,
			},
			ModuleDescription: sql.NullString{
				String: "<h2>JUDUL</h2>\n                    <p>Pembahasan tingkat menengah yang perlu diketahui tentang hal-hal apa saja yang harus dipelajari untuk membangun Enterpreneurship yang kaya akan adaptasi dan pembaharuan inovasi baru dalam membangun sebuah produk serta bermanfaat untuk masyarakat,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dipikat dari cara mengatasi berbagai rintangan dalam membangun produk selagi berkaya menuangkan semua kekreatifitasan dalam mendirikan nama dan produk yang baik untuk diterima masyarakat dan dapat membuat sebuah perubahan dimasa kini untuk kedepannya.</p>",
				Valid:  true,
			},
			IsPublished: true,
		},
		{
			ID:          6,
			CourseId:    5,
			ModuleTitle: "Title 2",
			ModuleOverview: sql.NullString{
				String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship.",
				Valid:  true,
			},
			ModuleDescription: sql.NullString{
				String: "<h2>JUDUL</h2>\n                    <p>Pembahasan tingkat menengah yang perlu diketahui tentang hal-hal apa saja yang harus dipelajari untuk membangun Enterpreneurship yang kaya akan adaptasi dan pembaharuan inovasi baru dalam membangun sebuah produk serta bermanfaat untuk masyarakat,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dipikat dari cara mengatasi berbagai rintangan dalam membangun produk selagi berkaya menuangkan semua kekreatifitasan dalam mendirikan nama dan produk yang baik untuk diterima masyarakat dan dapat membuat sebuah perubahan dimasa kini untuk kedepannya.</p>",
				Valid:  true,
			},
			IsPublished: true,
		},
		{
			ID:          7,
			CourseId:    5,
			ModuleTitle: "Title 3",
			ModuleOverview: sql.NullString{
				String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Pembahasan tingkat menengah yang perlu diketahui tentang hal-hal apa saja yang harus dipelajari untuk membangun Enterpreneurship yang kaya akan adaptasi dan pembaharuan inovasi baru dalam membangun sebuah produk serta bermanfaat untuk masyarakat,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dipikat dari cara mengatasi berbagai rintangan dalam membangun produk selagi berkaya menuangkan semua kekreatifitasan dalam mendirikan nama dan produk yang baik untuk diterima masyarakat dan dapat membuat sebuah perubahan dimasa kini untuk kedepannya.</p>"},
			IsPublished:       true,
		},
		{
			ID:                8,
			CourseId:          5,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Pembahasan tingkat menengah yang perlu diketahui tentang hal-hal apa saja yang harus dipelajari untuk membangun Enterpreneurship yang kaya akan adaptasi dan pembaharuan inovasi baru dalam membangun sebuah produk serta bermanfaat untuk masyarakat,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dipikat dari cara mengatasi berbagai rintangan dalam membangun produk selagi berkaya menuangkan semua kekreatifitasan dalam mendirikan nama dan produk yang baik untuk diterima masyarakat dan dapat membuat sebuah perubahan dimasa kini untuk kedepannya.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Expert to Entrepreneurship
		{
			ID:                9,
			CourseId:          9,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Pembahasan tingkat lanjut dengan para expert dan Orang-orang hebat yang membangun usaha dengan memberikan pengalaman yang berbobot serta menambah ilmu baru, juga membagikan beberapa tips dari pengalaman yang telah dialami setelah terjun langsung dalam membuat usaha yang bahkan membuat perubahan tak terduga,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>dikutip dari beberapa expert ini untuk meningkatkan esensi dan kelanjutan dalam bidang yang dipelajari khususnya dalam Enterpreneurship ini sehingga tidak hanya membuatkan suatu hal yang sia-sia melainkan orang-orang benar-benar mendapatkan hasil baru dari pembelajaran ini dan menerapkannya.</p>"},
			IsPublished:       true,
		},
		{
			ID:                10,
			CourseId:          9,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Pembahasan tingkat lanjut dengan para expert dan Orang-orang hebat yang membangun usaha dengan memberikan pengalaman yang berbobot serta menambah ilmu baru, juga membagikan beberapa tips dari pengalaman yang telah dialami setelah terjun langsung dalam membuat usaha yang bahkan membuat perubahan tak terduga,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>dikutip dari beberapa expert ini untuk meningkatkan esensi dan kelanjutan dalam bidang yang dipelajari khususnya dalam Enterpreneurship ini sehingga tidak hanya membuatkan suatu hal yang sia-sia melainkan orang-orang benar-benar mendapatkan hasil baru dari pembelajaran ini dan menerapkannya.</p>"},
			IsPublished:       true,
		},
		{
			ID:                11,
			CourseId:          9,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Pembahasan tingkat lanjut dengan para expert dan Orang-orang hebat yang membangun usaha dengan memberikan pengalaman yang berbobot serta menambah ilmu baru, juga membagikan beberapa tips dari pengalaman yang telah dialami setelah terjun langsung dalam membuat usaha yang bahkan membuat perubahan tak terduga,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>dikutip dari beberapa expert ini untuk meningkatkan esensi dan kelanjutan dalam bidang yang dipelajari khususnya dalam Enterpreneurship ini sehingga tidak hanya membuatkan suatu hal yang sia-sia melainkan orang-orang benar-benar mendapatkan hasil baru dari pembelajaran ini dan menerapkannya.</p>"},
			IsPublished:       true,
		},
		{
			ID:                12,
			CourseId:          9,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Enterpreneurship."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Pembahasan tingkat lanjut dengan para expert dan Orang-orang hebat yang membangun usaha dengan memberikan pengalaman yang berbobot serta menambah ilmu baru, juga membagikan beberapa tips dari pengalaman yang telah dialami setelah terjun langsung dalam membuat usaha yang bahkan membuat perubahan tak terduga,</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2070&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>dikutip dari beberapa expert ini untuk meningkatkan esensi dan kelanjutan dalam bidang yang dipelajari khususnya dalam Enterpreneurship ini sehingga tidak hanya membuatkan suatu hal yang sia-sia melainkan orang-orang benar-benar mendapatkan hasil baru dari pembelajaran ini dan menerapkannya.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Intro to Marketing
		{
			ID:                13,
			CourseId:          2,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari dasar-dasar dan pemahaman baru tentang marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                14,
			CourseId:          2,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari dasar-dasar dan pemahaman baru tentang marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                15,
			CourseId:          2,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari dasar-dasar dan pemahaman baru tentang marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                16,
			CourseId:          2,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari dasar-dasar dan pemahaman baru tentang marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Intermediate to Marketing
		{
			ID:                17,
			CourseId:          6,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari marketing tingkat menengah lanjutan dasar-dasar dari awal pemahaman dan fundamental marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                18,
			CourseId:          6,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari marketing tingkat menengah lanjutan dasar-dasar dari awal pemahaman dan fundamental marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                19,
			CourseId:          6,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari marketing tingkat menengah lanjutan dasar-dasar dari awal pemahaman dan fundamental marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                20,
			CourseId:          6,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari marketing tingkat menengah lanjutan dasar-dasar dari awal pemahaman dan fundamental marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Expert to Marketing
		{
			ID:                21,
			CourseId:          10,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari marketing tingkat menengah lanjutan dasar-dasar dari awal pemahaman dan fundamental marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                22,
			CourseId:          10,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari marketing tingkat menengah lanjutan dasar-dasar dari awal pemahaman dan fundamental marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                23,
			CourseId:          10,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari marketing tingkat menengah lanjutan dasar-dasar dari awal pemahaman dan fundamental marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		{
			ID:                24,
			CourseId:          10,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari marketing tingkat menengah lanjutan dasar-dasar dari awal pemahaman dan fundamental marketing. memiliki pengetahuan dasar dan membangun fundamental yang harus dimiliki oleh orang yang berjiwa marketing, bagaimana memahami pelanggan dan masyarakat tempat beradu untuk menyajikan produk-produk yang ingin dijual serta inovatif. kerjasama serta disiplin serta hal-hal teknis lain seperti bagaimana cara berbicara menyajikan iklan dan lain-lain</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1552664730-d307ca884978?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Sesuai dengan keutamaan kelompok dan juga bidang-bidang yang ingin dikuasai dengan pengalaman yang ada pada marketing banyak hal yang bisa dijadikan referensi baru dan pengetahuan baru yang bisa diambil.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Intro to English
		{
			ID:                25,
			CourseId:          3,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                26,
			CourseId:          3,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                27,
			CourseId:          3,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                28,
			CourseId:          3,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Intermediate to English
		{
			ID:                29,
			CourseId:          7,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                30,
			CourseId:          7,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                31,
			CourseId:          7,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                32,
			CourseId:          7,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Expert to English
		{
			ID:                33,
			CourseId:          11,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                34,
			CourseId:          11,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                35,
			CourseId:          11,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		{
			ID:                36,
			CourseId:          11,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Marketing."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mempelajari fundamental dan dasar-dasar dalam berbahasa inggris. belajar bagaimana pengucapan dalam berbicara, membaca, grammar, kosakata baru, pendengaran, serta penulisan dalam berbahasa inggris, agar memiliki pengalaman juga dalam hal menggunakan bahasa inggris sebagai bahasa yang dapat digunakan sehari-hari dan berkomunikasi dengan orang lain.</p>\n                    <br>\n                    <p><img src=\"https://images.unsplash.com/photo-1453748866136-b1dd97284f49?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=818&q=80\" alt=\"Testing Photo\"></p>\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>Dengan pengalaman baru dan pengetahuan baru menambahkan beberapa kesan dan juga ilmu yang dapat berguna dalam penggunaan bahasa inggris juga agar dapat di implementasikan bukan hanya sekedar lewat, tetapi sungguh dasar-dasarnya dapat dimiliki dan kelancaran dalam bebahasa inggris.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Intro to Creative
		{
			ID:                37,
			CourseId:          4,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                38,
			CourseId:          4,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                39,
			CourseId:          4,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                40,
			CourseId:          4,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Intermediate to Creative
		{
			ID:                41,
			CourseId:          8,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                42,
			CourseId:          8,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                43,
			CourseId:          8,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                44,
			CourseId:          8,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		//#

		//#Module Expert to Creative
		{
			ID:                45,
			CourseId:          12,
			ModuleTitle:       "Title 1",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                46,
			CourseId:          12,
			ModuleTitle:       "Title 2",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                47,
			CourseId:          12,
			ModuleTitle:       "Title 3",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		{
			ID:                48,
			CourseId:          12,
			ModuleTitle:       "Title 4",
			ModuleOverview:    sql.NullString{String: "Alami Pembelajaran langsung dari para expertnya dan memiliki pemahaman baru tentang Creative."},
			ModuleDescription: sql.NullString{String: "<h2>JUDUL</h2>\n                    <p>Mengenal dan mempelajari dasar-dasar serta membangun fundamental baru dalam memiliki kekreatifitasan yang dapat dibangun juga dari pembelajaran-pembelajaran serta pengetahuan yang memperluas cara pandang dan kreatifitaspun dapat terbentuk dan dapat digunakan untuk banyak hal, yang mana kreatifitas sudah menjadi salah satu hal yang sangat penting untuk dimiliki dalam dunia kerja bahkan membangun usaha.</p>\n                    <br>\n                    <img src=\"https://images.unsplash.com/photo-1560421683-6856ea585c78?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=874&q=80\" alt=\"Testing Photo\">\n                    <br>\n                    <h3>SUB JUDUL</h3>\n                    <p>kreatifitas menjadi kunci baru dalam membuka banyak hal, salah satu contohnya membuka peluang kerja baru, dengan adanya kreatifitas membangun sesuatu yang baru dan invoatif demi perkembangan zaman pun bisa terjadi karena merupakasan suatu kunci kesuksesan dan pembaharuan bagi masyarakat dan dunia dalam menciptakan sesuatu yang belum pernah ditemukan.</p>"},
			IsPublished:       true,
		},
		//#
	}
}

func (s *MasterElearningModuleSeeder) Seed(db *gorm.DB) error {
	elearningModules := InitElearningModuleData()

	return db.CreateInBatches(elearningModules, len(elearningModules)).Error
}

func (s *MasterElearningModuleSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterElearningModule{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
