name 'dappdeps-base'

license 'MIT'
license_file 'https://github.com/flant/dappdeps-base/blob/master/LICENSE.txt'

dependency 'bash'
dependency 'gtar'
dependency 'sudo'
dependency 'findutils'
dependency 'sed'

build do
  link "#{install_dir}/embedded/bin", "#{install_dir}/bin"
end
