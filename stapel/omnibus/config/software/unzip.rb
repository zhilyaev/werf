name "unzip"
default_version "6.0"

version("6.0") { source md5: "62b490407489521db863b523a7f86375" }

license "Info-ZIP"
license_file "LICENSE"

source url: "https://downloads.sourceforge.net/infozip/unzip60.tar.gz"

relative_path "unzip60"

build do
  env = with_standard_compiler_flags(with_embedded_path)

  configure_command = [
    "./configure",
    "--prefix=#{install_dir}/embedded",
  ]

  command "./configure" \
          " --prefix=#{install_dir}/embedded", env: env

  make "-j #{workers}", env: env
  make "-j #{workers} install", env: env
end
