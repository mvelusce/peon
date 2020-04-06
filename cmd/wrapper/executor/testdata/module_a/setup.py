"""setup module"""

import setuptools

setuptools.setup(name='module_a',
                 version='1.2.3',
                 description='Module a',
                 url='some-repo',
                 author='sky',
                 zip_safe=False,
                 long_description_content_type='text/markdown',
                 packages=setuptools.find_packages(),
                 include_package_data=True,
                 test_suite='tests',
                 scripts=['bin/run_module_a.py'])
